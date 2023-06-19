<template>
  <main class="home">
    <div>
      {{ t("home.welcome") }}{{store.user.name}}
    </div>
  </main>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useStore } from "@/stores/store";

const { t } = useI18n();
const store = useStore();

</script>


<style scoped lang="scss">

</style>
