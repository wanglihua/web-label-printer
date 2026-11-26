<script setup>
defineOptions({
    name: "page-not-found",
});
</script>

<template>
    <div class="main">
        <el-result
            icon="error"
            title="Page Not Found"
            sub-title="The page you visited does not exist. Please enter a valid web address"
        >
            <template #extra>
                <el-button type="primary" @click="() => {
                if (this.$route.fullPath.startsWith('/admin')) {
                    this.$router.push('/admin');
                }
                else
                {
                    this.$router.push('/');
                }
            }">Home
                </el-button>
            </template>
        </el-result>
    </div>
</template>

<style scoped>
.main {
    padding-top: 120px;
}
</style>