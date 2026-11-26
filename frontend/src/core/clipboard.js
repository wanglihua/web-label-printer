const copyText = (text) => {
    const ta = document.createElement('textarea');
    ta.value = text;
    ta.cols = 1;
    ta.rows = 1;
    ta.style.color = 'transparent';
    ta.style.border = 'none';
    document.body.appendChild(ta);
    ta.select();
    let success = false;
    try {
        success = document.execCommand('copy');
    } catch (err) {
        success = false;
    }
    document.body.removeChild(ta);
    return success;
};

export default copyText;
