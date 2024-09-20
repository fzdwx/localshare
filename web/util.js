// 生成UUID的函数
function generateUUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

// 获取或创建用户ID
function getUserId() {
    let id = localStorage.getItem('chatUserId');
    if (!id) {
        id = generateUUID();
        localStorage.setItem('chatUserId', id);
    }
    return id;
}
