import {reactive} from 'vue';

// Global shared state and some configuration values
const store = reactive({
    navMenu: "",
    navMenuCollapsed: false,

    messageList: [],
    loadMessageList() {
        // Fill in this call logic

        // notify notification message content
    },

    paginationPageSizes: [10, 15, 20, 30, 50, 100],
    paginationDefaultPageSize: 15,
    paginationLayout: "total, sizes, prev, pager, next, jumper",

    // task home
});

export {store};