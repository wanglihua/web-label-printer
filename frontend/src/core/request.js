import axios from "axios";
import {hideMask, showMask} from "./mask.js";
import {showError} from "./message.js";

const post = (url, data, fn) => {
    showMask();
    axios.request({
        url: url,
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        responseType: 'json',
        responseEncoding: 'utf8',
        data: data,
        transformRequest: [function (data) {
            return JSON.stringify(data);
        }]
    }).then((response) => {
        hideMask();

        if (response.data.success !== true && response.data.needLogin === true) {
            showError(response.data.message, "Login and Permissions", function () {
                window.location.href = "/login";
            });
            return;
        }

        if (response.data.success !== true && response.data.needTaskSelect === true) {
            showError(response.data.message, "Task Selection", function () {
                window.location.href = "/task/select";
            });
            return;
        }

        if (response.data.success !== true) {
            showError(response.data.message);
            return;
        }

        if (fn) {
            fn(response.data);
        }
    }).catch((error) => {
        hideMask();
        // Usually a URL error, network error, etc.
        showError(`An error occurred while processing the request. Please refresh and try again ${error.toString()}`);
    });
};

const postNoErrorHandle = (url, data, fn) => {
    showMask();
    axios.request({
        url: url,
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        responseType: 'json',
        responseEncoding: 'utf8',
        data: data,
        transformRequest: [function (data) {
            return JSON.stringify(data);
        }]
    }).then((response) => {
        hideMask();

        if (fn) {
            fn(response.data);
        }

    }).catch((error) => {
        hideMask();
        // Usually a URL error, network error, etc.
        showError(`An error occurred while processing the request. Please refresh and try again ${error.toString()}`);
    });
};

const postNoMask = (url, data, fn) => {
    axios.request({
        url: url,
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        responseType: 'json',
        responseEncoding: 'utf8',
        data: data,
        transformRequest: [function (data) {
            return JSON.stringify(data);
        }]
    }).then((response) => {
        if (response.data.success !== true && response.data.needLogin === true) {
            showError(response.data.message, "Login and Permissions", function () {
                window.location.href = "/login";
            });
            return;
        }

        if (response.data.success !== true && response.data.needTaskSelect === true) {
            showError(response.data.message, "Task Selection", function () {
                window.location.href = "/task/select";
            });
            return;
        }

        if (response.data.success !== true) {
            showError(response.data.message);
            return;
        }

        if (fn) {
            fn(response.data);
        }
    }).catch((error) => {
        // Usually a URL error, network error, etc.
        showError(`An error occurred while processing the request. Please refresh and try again ${error.toString()}`);
    });
};

const postNoErrorHandleNoMask = (url, data, fn) => {
    axios.request({
        url: url,
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        responseType: 'json',
        responseEncoding: 'utf8',
        data: data,
        transformRequest: [function (data) {
            return JSON.stringify(data);
        }]
    }).then((response) => {
        if (fn) {
            fn(response.data);
        }
    }).catch((error) => {
        // Usually a URL error, network error, etc.
        showError(`An error occurred while processing the request. Please refresh and try again ${error.toString()}`);
    });
};

export {post, postNoErrorHandle, postNoMask, postNoErrorHandleNoMask};
