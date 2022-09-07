import {
    createApp as createClientApp,
    h,
    VNode,
    resolveDynamicComponent,
    Transition,
    App,
    onMounted,
} from 'vue';
import { RouterView } from 'vue-router';
import { createPinia } from 'pinia';
import { createRouter } from './router';

//import VReloadPrompt from './components/base/modal/VReloadPrompt.vue';

type VueroAppOptions = {
    enhanceApp?: (app: App) => Promise<void>;
};

export async function createApp({ enhanceApp }: VueroAppOptions) {
    const router = createRouter();
    const pinia = createPinia();

    const app = createClientApp({
        // This is the global app setup function
        setup() {

            /**
             * Here you can check if your user has a token stored
             * and check with your api if it still valid before your app start
             */
            onMounted(async () => {

            });
            return () => {
                const defaultSlot = ({ Component: _Component }: any) => {
                    const Component = resolveDynamicComponent(_Component) as VNode;

                    return [
                        h(
                            Transition,
                            { name: 'fade-slow', mode: 'out-in' },
                            {
                                default: () => [h(Component)],
                            }
                        ),
                    ];
                };

                return [
                    h(RouterView, null, {
                        default: defaultSlot,
                    }),
                    //h(VReloadPrompt, { appName: 'MongoViewer' }),
                ];
            };
        },
    });

    app.use(router);
    app.use(pinia);

    if (enhanceApp) {
        await enhanceApp(app);
    }

    return {
        app,
        router,
    };
}