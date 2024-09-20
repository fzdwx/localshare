const chatContainer = document.getElementById('chat-container');
const messageInput = document.getElementById('message-input');
const fileInput = document.getElementById('file-input');
const sendButton = document.getElementById('send-button');
const body = document.body;

let ws;
let userId;
let pasteFile = null;

// 初始化用户ID
userId = getUserId();
console.log('User ID:', userId);

function connectWebSocket() {
    let url = window.location.host + '/ws';
    ws = new WebSocket(`ws://${url}`);

    ws.onopen = () => {
        console.log('WebSocket连接已打开');
        // 连接成功后发送用户ID
        ws.send(JSON.stringify({type: 'identify', sender: userId}));
    };
    ws.onclose = () => {
        console.log('WebSocket连接已关闭');
        setTimeout(connectWebSocket, 3000); // 3秒后尝试重新连接
    };

    ws.onmessage = (event) => {
        if (event.data instanceof Blob) {
            handleBinaryMessage(event.data);
        } else {
            const message = JSON.parse(event.data);
            appendMessage(message);
        }
    };
}

connectWebSocket();

function handleBinaryMessage(blob) {
    const reader = new FileReader();
    reader.onload = (e) => {
        const message = {
            type: 'file',
            fileName: blob.name || 'unknown',
            fileType: blob.type,
            fileContent: e.target.result,
            sender: 'other'
        };
        appendMessage(message);
    };
    reader.readAsDataURL(blob);
}

function appendMessage(message) {
    const messageElement = document.createElement('div');
    messageElement.classList.add('flex', message.sender === userId ? 'justify-end' : 'justify-start');

    const contentElement = document.createElement('div');
    contentElement.classList.add('max-w-xs', 'lg:max-w-md', 'rounded-lg', 'p-3', 'shadow');

    if (message.sender === userId) {
        contentElement.classList.add('bg-indigo-500', 'text-white');
    } else {
        contentElement.classList.add('bg-gray-100', 'text-gray-800');
    }

    if (message.type === 'identify') {
        contentElement.textContent = `User ${message.sender} joined the chat`;
        contentElement.classList = ['text-center', 'text-gray-500', 'text-sm'];
    } else if (message.type === 'text') {
        const urlRegex = /(https?:\/\/[^\s]+)/g;
        const parts = message.text.split(urlRegex);
        parts.forEach(part => {
            if (urlRegex.test(part)) {
                const link = document.createElement('a');
                link.href = part;
                link.textContent = part;
                link.classList.add('text-blue-600', 'underline');
                link.target = '_blank';
                contentElement.appendChild(link);
            } else {
                const textNode = document.createTextNode(part);
                contentElement.appendChild(textNode);
            }
        });
    } else if (message.type === 'file') {
        if (message.fileType.startsWith('image/')) {
            const img = document.createElement('img');
            img.src = message.fileContent;
            img.classList.add('max-w-full', 'h-auto', 'rounded');
            contentElement.appendChild(img);
        } else {
            const link = document.createElement('a');
            link.href = message.fileContent;
            link.download = message.fileName;
            link.textContent = `Download ${message.fileName}`;
            link.classList.add('text-blue-600', 'underline');
            contentElement.appendChild(link);
        }
    }

    messageElement.appendChild(contentElement);
    chatContainer.appendChild(messageElement);
    window.scrollTo(0, body.scrollHeight)
}

function sendMessage() {
    const messageText = messageInput.value.trim();
    let file = pasteFile
    if (file === null) {
        file = fileInput.files[0];
    }

    if ((!messageText && !file) || !ws || ws.readyState !== WebSocket.OPEN) return;

    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
            const fileContent = e.target.result;
            const message = {
                type: 'file',
                fileName: file.name,
                fileType: file.type,
                fileContent: fileContent,
                sender: userId
            };
            ws.send(JSON.stringify(message));
            appendMessage(message);
            fileInput.value = '';
        };
        reader.readAsDataURL(file);
    }

    if (messageText) {
        const message = {type: 'text', text: messageText, sender: userId};
        ws.send(JSON.stringify(message));
        appendMessage(message);
        messageInput.value = '';
    }
}

sendButton.addEventListener('click', sendMessage);
messageInput.addEventListener('keypress', (event) => {
    if (event.key === 'Enter') {
        sendMessage();
    }
});

messageInput.addEventListener('paste', (event) => {
    let file = null;
    const items = (event.clipboardData || window.clipboardData).items;
    if (items && items.length) {
        for (var i = 0; i < items.length; i++) {
            if (items[i].type.includes('image')) {
                file = items[i].getAsFile();
                break;
            }
        }
    }
    if (file) {
        pasteFile = file;
        sendMessage()
        pasteFile = null;
    }
});

fileInput.addEventListener('change', () => {
    if (fileInput.files.length > 0) {
        sendMessage();
    }
});