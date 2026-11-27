<script setup>
import {onMounted, ref} from "vue";
import PageHeader from "../page-header.vue";
import {post} from "../../core/prelude.js";

defineOptions({
    name: "admin-home",
});

onMounted(() => {
    post("/api/admin/home_data", {}, function (res) {
    });
});
</script>

<template>
    <page-header menu="admin_home"/>
</template>

<style scoped>

</style>
