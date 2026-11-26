import {markRaw} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import IconInfo from "./icon-info.vue";
import IconQuestion from "./icon-question.vue";
import IconSuccess from "./icon-success.vue";

const showSuccess = (message) => {
    /*

    import {ElNotification} from "element-plus"

    if (!title) {
        title = "Success";
    }

    ElNotification({
        title: title,
        message: message,
        type: "success",
    });

     */


    ElMessage({
        message: message,
        type: "success",
    });
};

const showError = (message, title, fn) => {
    if (!title) {
        title = "Error";
    }

    ElMessageBox.alert(
        message,
        title,
        {
            confirmButtonText: "Confirm",
            showCancelButton: false,
            type: "error",
            draggable: true,
            // center: true,
            closeOnClickModal: false,
            // buttonSize: "large",
        }
    ).then(() => {
        if (fn) {
            fn();
        }
    });
};

const showConfirm = (message, fn, title) => {
    if (!title) {
        title = "Confirm";
    }

    ElMessageBox.confirm(
        message,
        title,
        {
            confirmButtonText: "Confirm",
            cancelButtonText: "Cancel",
            type: "info",
            draggable: true,
            closeOnClickModal: false,
            icon: markRaw(IconQuestion),
        }
    ).then(() => {
        if (fn) {
            fn();
        }
    }).catch(() => {

    })
};

const showInfo = (message, fn, title) => {
    if (!title) {
        title = "Notice";
    }

    ElMessageBox.alert(message, title, {
        confirmButtonText: "Confirm",
        type: 'info',
        draggable: true,
        closeOnClickModal: false,
        icon: markRaw(IconInfo),
        callback: () => {
            if (fn) {
                fn();
            }
        },
    }).then(() => {
    }).catch(() => {
    });
};

const showInfoSuccess = (message, fn, title) => {
    if (!title) {
        title = "Success";
    }

    ElMessageBox.alert(message, title, {
        confirmButtonText: "Confirm",
        type: 'info',
        draggable: true,
        closeOnClickModal: false,
        icon: markRaw(IconSuccess),
        callback: () => {
            if (fn) {
                fn();
            }
        },
    }).then(() => {
    }).catch(() => {
    });
};

export {showConfirm, showError, showInfo, showInfoSuccess, showSuccess};