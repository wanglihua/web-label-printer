import axios from "axios";
import {hideMask, showMask} from "./mask.js";
import {showError} from "./message.js";

const download = (url, data) => {
    if (!data) {
        data = {};
    }

    showMask();
    axios.request({
        url: url,
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        responseType: 'blob', // important
        data: data,
        transformRequest: [function (data) {
            return JSON.stringify(data);
        }]
    }).then((response) => {
        hideMask();

        let headerAttachment = response.headers["content-disposition"]; // attachment; filename=___Product.xlsx; filename*=UTF-8''%E4%BA%A7%E5%93%81_Product.xlsx
        if (!headerAttachment) {
            const reader = new FileReader();
            reader.onload = function (event) {
                const data = JSON.parse(event.target.result);

                if (data.success === false && data.needLogin === true) {
                    showError(data.message, "Login and Permissions", function () {
                        window.location.href = "/login";
                    });
                    return;
                }

                if (data.success === false) {
                    showError(data.message);
                    return;
                }

                showError("Download failed. An unknown error occurred");
            };
            reader.readAsText(response.data);

            return;
        }

        // ge file name from response header
        let fileNamePrefix = "filename*=UTF-8''";
        let fileNamePrefixPos = headerAttachment.lastIndexOf(fileNamePrefix);
        let fileNamePrefixLen = fileNamePrefix.length;
        let fileName = decodeURIComponent(headerAttachment.substring(fileNamePrefixPos + fileNamePrefixLen));

        // create file link in browser's memory
        const href = URL.createObjectURL(response.data);

        // create "a" HTML element with href to file & click
        const link = document.createElement('a');
        link.href = href;
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();

        // clean up "a" element & remove ObjectURL
        document.body.removeChild(link);
        URL.revokeObjectURL(href);
    }).catch((error) => {
        hideMask();
        // Usually a URL error, network error, etc.
        showError(`An error occurred while processing the request. Please refresh and try again ${error.toString()}`);
    });
};

export default download;