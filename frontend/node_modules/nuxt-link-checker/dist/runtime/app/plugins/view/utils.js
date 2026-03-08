import { getCurrentScope, onScopeDispose } from "vue";
export function useEventListener(target, type, listener, options) {
  target.addEventListener(type, listener, options);
  if (getCurrentScope()) {
    onScopeDispose(() => target.removeEventListener(type, listener, options));
  }
}
