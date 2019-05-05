import Vue from 'vue';
import isFunction from 'lodash/isFunction';
let handleOutsideClick;
Vue.directive('closable', {
    bind(el, binding, vnode) {
        handleOutsideClick = e => {
            e.stopPropagation();
            const {handler, exclude = []} = binding.value;
            let clickedOnExcludedEl = false;
            exclude.forEach(refName => {
                if (!clickedOnExcludedEl) {
                    const excludedEl = vnode.context.$refs[refName];
                    clickedOnExcludedEl = excludedEl.contains(e.target);
                }
            });
            if (!el.contains(e.target) && !clickedOnExcludedEl) {
                if (isFunction(handler)) {
                    handler();
                }
            }
        };
        // Register click/touchstart event listeners on the whole page
        document.addEventListener('click', handleOutsideClick);
        document.addEventListener('focusin', handleOutsideClick);
        document.addEventListener('touchstart', handleOutsideClick);
    },
    unbind() {
        document.removeEventListener('click', handleOutsideClick);
        document.removeEventListener('focusin', handleOutsideClick);
        document.removeEventListener('touchstart', handleOutsideClick);
    }
});
