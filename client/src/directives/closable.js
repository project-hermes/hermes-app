import Vue from 'vue';
import isFunction from 'lodash/isFunction';
let handleOutsideClick;
Vue.directive('closable', {
    bind(el, binding) {
        handleOutsideClick = e => {
            e.stopPropagation();
            // const {handler, exclude = []} = binding.value;
            let clickedOnExcludedEl = false;
            // exclude.forEach(refName => {
            //     if (!clickedOnExcludedEl) {
            //         const excludedEl = vnode.context.$refs[refName];
            //         clickedOnExcludedEl = excludedEl.contains(e.target);
            //     }
            // });
            if (!el.contains(e.target) && !clickedOnExcludedEl) {
                // vnode.context[handler]();
                if (isFunction(binding.value)) {
                    binding.value();
                }
            }
        };
        // Register click/touchstart event listeners on the whole page
        document.addEventListener('click', handleOutsideClick);
        document.addEventListener('touchstart', handleOutsideClick);
    },
    unbind() {
        document.removeEventListener('click', handleOutsideClick);
        document.removeEventListener('touchstart', handleOutsideClick);
    }
});
