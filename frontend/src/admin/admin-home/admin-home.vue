<script setup>
import { onMounted, ref } from "vue";
import PageHeader from "../page-header.vue";
import { post, showSuccess } from "../../core/prelude.js";

defineOptions({
    name: "admin-home",
});

onMounted(() => {
    post("/api/admin/home_data", {}, function (res) {
    });
});

function printTest() {
    post("/api/admin/print_test", {}, function (res) {
        const pdfUrl = res.pdfUrl;

        post("http://localhost:8113/print", {
            pdf: pdfUrl
        }, function (res) {
            showSuccess(res.message);
        });
    });
}

</script>

<template>
    <page-header menu="admin_home" />
    <div class="page-main">
        <el-button type="primary" @click="printTest">Print Test</el-button>
    </div>
</template>

<style scoped></style>
