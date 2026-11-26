<script setup>
import {useRouter} from "vue-router";
import {store} from '../core/prelude.js';
import {HomeFilled} from "@element-plus/icons-vue";

defineOptions({
    name: 'admin-layout',
});

const router = useRouter();

</script>

<template>
    <el-container class="admin-layout-container">
        <el-header class="admin-layout-header">
            <RouterLink to="/admin" style="text-decoration: none;">
                <div class="title">
                    <div class="logo"></div>
                    <div class="text">printer</div>
                </div>
            </RouterLink>
        </el-header>
        <el-container>
            <el-aside class="admin-layout-aside" :style="{width: store.navMenuCollapsed === true? '60px': '180px'}">
                <el-menu
                        :default-active="store.navMenu"
                        :unique-opened="true"
                        :collapse="store.navMenuCollapsed"
                        :router="true"
                >
                    <el-menu-item index="admin_home" route="/admin">
                        <el-icon><HomeFilled /></el-icon>
                        <span>Home</span>
                    </el-menu-item>
                </el-menu>
            </el-aside>
            <el-container class="admin-layout-main-container">
                <el-main>
                    <div class="admin-main-content">
                        <RouterView/>
                    </div>
                </el-main>
                <el-footer></el-footer>
            </el-container>
        </el-container>
    </el-container>
</template>

<style scoped>
.admin-layout-container {
    min-width: 860px;
    height: 100vh;
}

.admin-layout-header {
    height: 55px;
    line-height: 55px;
    padding: 0;
    display: -webkit-flex;
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-between;
    background-color: #4285f4;
    color: white;
    /*border-bottom: 1px solid var(--el-border-color);*/
    box-shadow: var(--el-box-shadow-light);
}

.admin-layout-header div.title {
    display: -webkit-flex;
    display: flex;
    flex-flow: row nowrap;
}

.admin-layout-header div.logo {
    margin-left: 5px;
    height: 55px;
    width: 55px;
    background-image: url("../assets/images/logo.webp");
    background-repeat: no-repeat;
    background-position: center center;
    cursor: pointer;
}

.admin-layout-header div.text {
    padding-left: 10px;
    font-size: 24px;
    text-decoration: none;
    color: white;
}

.admin-layout-header div.user {
    padding-right: 10px;
    font-size: 14px;
    cursor: pointer;
}

.admin-layout-header div.user .iconfont {
    font-size: 18px;
    padding-left: 8px;
}

.admin-layout-aside {
    --el-menu-bg-color: var(--admin-main-bg-color);
    --el-menu-active-color: var(--el-color-primary);
    --el-menu-text-color: var(--el-text-color-primary);
    --el-menu-hover-text-color: var(--el-color-primary);
    --el-menu-hover-bg-color: var(--el-color-primary-light-9);
    --el-menu-item-height: 48px;
    --el-menu-sub-item-height: calc(var(--el-menu-item-height) - 4px);
    --el-menu-item-font-size: var(--el-font-size-base);
    --el-menu-item-hover-fill: var(--el-color-primary-light-9);
    --el-menu-border-color: var(--el-menu-bg-color);
    --el-menu-base-level-padding: 16px;
    --el-menu-level-padding: 20px;
    --el-menu-icon-width: 24px;

    background-color: var(--el-menu-bg-color);
}

.admin-layout-aside .el-sub-menu.is-active.is-opened {
    --el-menu-bg-color: #fbfbfb;
}

.admin-layout-aside .el-menu-item.is-active {
    background-color: var(--el-menu-hover-bg-color)
}

.admin-layout-aside .iconfont {
    font-size: 18px;
    padding-right: 6px;
    margin-top: -4px;
}

.admin-layout-main-container .el-main {
    --el-main-padding: 0;
}

.admin-nav-menu-link {
    display: block;
    width: 165px;
    margin-left: -5px;
    padding-left: 5px;
    text-decoration: none;
    color: var(--el-menu-text-color);
}

</style>