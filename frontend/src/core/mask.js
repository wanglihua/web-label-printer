import {ElLoading} from "element-plus";

let loading;

const showMask = () => {
    loading = ElLoading.service({
        lock: false,
        text: "",
        background: 'rgba(0, 0, 0, 0.05)',
    });
};

const hideMask = () => {
    if (loading) {
        loading.close();
    }
};

export {showMask, hideMask};