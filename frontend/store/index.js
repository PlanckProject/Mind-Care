import Vuex from 'vuex';
import theme from '~/store/modules/theme.js';

const createStore = () => {
    return new Vuex.Store({
        modules: {
            theme
        }
    })
}

export default createStore;