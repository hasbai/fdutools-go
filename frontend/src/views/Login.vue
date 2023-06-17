<template>
  <main class="flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full">
      <h2 class="text-3xl font-semibold mb-6 text-center">{{ t("login.login") }}</h2>
      <form class="bg-white shadow-md rounded px-8 py-6 mb-4" @submit.prevent="login">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="username">
            {{ t("login.username") }}
          </label>
          <input
            v-model="username"
            :class="{'border-red-500': showError}"
            class="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            type="text"
            :placeholder="t('login.username')"
            required
          />
        </div>
        <div class="mb-2">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
            {{ t("login.password") }}
          </label>
          <input
            v-model="password"
            :class="{'border-red-500': showError}"
            class="appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
            type="password"
            :placeholder="t('login.password')"
            required
          />
        </div>
        <div class="flex flex-col items-center justify-between">
          <button
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            {{ t("login.login") }}
          </button>
          <p v-if="showError" class="text-center mt-4" :class="{'text-red-600': showError}">
            {{error}}
          </p>
        </div>
      </form>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Login } from "../../wailsjs/go/fdu/Fdu";
import { useStore } from "@/stores/store";
import router from "@/router";

import { useI18n } from "vue-i18n";
const { t } = useI18n();

const username = ref('');
const password = ref('');
const showError = ref(false);
const error = ref('');

const store = useStore();
async function login() {
  showError.value = false
  try{
    const user = await Login([username.value, password.value]);
    store.setUser(user)
    await router.push("/");
  } catch (e: any) {
    showError.value = true;
    error.value = e
  }
}

</script>

<style>
/* Add any custom styles here */
</style>
