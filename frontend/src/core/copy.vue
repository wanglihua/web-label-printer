<script setup>
import {showSuccess} from "./message.js";
import copyText from "./clipboard.js";
import {CopyDocument} from "@element-plus/icons-vue";

defineOptions({
    name: "copy",
});

const props = defineProps({
    text: String,
});

function textClick() {
    copyText(props.text);
    showSuccess(`Copied successfully: ${props.text}`);
}
</script>

<template>
    <span style="cursor: pointer; padding-right: 16px" @click="textClick">
        <el-icon style="color: var(--el-color-primary-light-7); position: relative; top: 2px; left: 2px">
            <copy-document/>
        </el-icon>
    </span>
</template>

<style scoped></style>
