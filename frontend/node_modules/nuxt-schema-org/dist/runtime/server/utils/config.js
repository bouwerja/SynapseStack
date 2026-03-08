import { defu } from "defu";
import { useRuntimeConfig } from "nitropack/runtime";
export function useSchemaOrgConfig(e) {
  const runtimeConfig = useRuntimeConfig(e);
  return defu(import.meta.client ? runtimeConfig.public["nuxt-schema-org"] : runtimeConfig["nuxt-schema-org"], {
    scriptAttributes: {}
  });
}
