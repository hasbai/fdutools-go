<template>
  <Transition name="fade">
    <div
      v-show="show"
      class="modal fixed inset-0 flex items-center justify-center z-50"
    >
      <div
        class="modal-overlay absolute inset-0 bg-gray-500 opacity-75"
        @click="close"
      ></div>
      <div class="modal-container bg-white w-96 mx-auto rounded shadow-lg z-50">
        <div class="modal-content flex flex-col items-center px-4 py-6">
          <h2 class="text-lg font-semibold">
            {{ store.currentCourse.name }}
          </h2>

          <div class="my-6">
            <input
              type="checkbox"
              class="sr-only"
              id="xk-option-checkbox"
              :checked="skipCaptcha"
              @change="skipCaptcha = !skipCaptcha"
            />
            <label class="flex items-center" for="xk-option-checkbox">
              <span class="mr-4">{{ t("xk.skip-captcha") }}</span>
              <span
                class="relative cursor-pointer w-10 h-6 bg-gray-300 rounded-full shadow-inner transition duration-300 ease-in-out"
              >
                <span
                  class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full shadow transition duration-300 ease-in-out transform"
                  :class="{
                    'translate-x-4': skipCaptcha,
                    'translate-x-0': !skipCaptcha,
                  }"
                ></span>
              </span>
            </label>
          </div>

          <div class="flex justify-evenly w-full">
            <button
              @click="drop"
              class="px-4 py-1 text-white rounded bg-red-700 hover:bg-red-600"
            >
              {{ t("xk.drop") }}
            </button>
            <button
              @click="select"
              class="px-4 py-1 text-white rounded bg-green-700 hover:bg-green-600"
            >
              {{ t("xk.select") }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { useStore } from "@/stores/store";
import { computed, ref } from "vue";
import { xk } from "../../wailsjs/go/models";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const store = useStore();

const show = computed(() => {
  return Boolean(store.currentCourse.id);
});

const skipCaptcha = ref(false);

const select = () => {
  close();
};

const drop = () => {
  close();
};

const close = () => {
  store.setCurrentCourse(new xk.Course());
};
</script>

<style scoped>
input:checked + label > span:nth-child(2) {
  background-color: #10b981;
}
</style>
