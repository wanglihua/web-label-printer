<script setup>
import {onMounted} from "vue";
import {store} from "../core/prelude.js";

defineOptions({
    name: "page-header",
});

const props = defineProps({
    menu: String,
});

onMounted(() => {
    store.navMenu = props.menu;
});

function toggleNavMenuCollapse() {
    store.navMenuCollapsed = store.navMenuCollapsed !== true;
}
</script>

<template>
    <div class="page-header">
        <div style="display: flex;">
            <div style="padding-right: 10px; cursor: pointer;" @click="toggleNavMenuCollapse">
                <el-icon v-if="store.navMenuCollapsed === true">
                    <Expand/>
                </el-icon>
                <el-icon v-else>
                    <Fold/>
                </el-icon>
            </div>
            <div>
                <el-breadcrumb separator="/">
                    <el-breadcrumb-item to="/admin"><i class="iconfont icon-home"/></el-breadcrumb-item>
                    <slot></slot>
                </el-breadcrumb>
            </div>
        </div>
    </div>
</template>

<style scoped>

.page-header {
    padding: 15px 20px;
    background-color: var(--admin-main-bg-color);
    /* border-left: 1px solid var(--el-border-color-light); */
}

</style>