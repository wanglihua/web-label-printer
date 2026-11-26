import {createApp} from 'vue';
import ElementPlus from 'element-plus';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import IconInfo from "./core/icon-info.vue";
import IconQuestion from "./core/icon-question.vue";
import IconSuccess from "./core/icon-success.vue";
import 'element-plus/dist/index.css';
import './index.css';
import App from './app.vue';
import router from './router';

const app = createApp(App);
app.config.errorHandler = (err, instance, info) => {
    // error - the thrown error object
    // instance - the component instance that threw the error
    // info - error code - https://vuejs.org/error-reference/

    console.log(err);
    console.log(instance);
    console.log(info);
}

app.use(ElementPlus);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

app.component("IconInfo", IconInfo);
app.component("IconQuestion", IconQuestion);
app.component("IconSuccess", IconSuccess);

app.use(router).mount('#root');