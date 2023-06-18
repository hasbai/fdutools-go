<template>
  <header>
    <div class="nav">
      <router-link to="/">{{ t("nav.home") }}</router-link>
      <router-link to="/grade">{{ t("nav.grade") }}</router-link>
      <router-link to="/login">{{ t("nav.login") }}</router-link>
      <router-link to="/about">{{ t("nav.about") }}</router-link>
    </div>
    <div class="menu">
      <div class="language">
        <div
          v-for="item in languages"
          :key="item"
          :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)"
          class="lang-item"
        >
          {{ t("languages." + item) }}
        </div>
      </div>
    </div>
  </header>
  <router-view />
</template>

<script setup lang="ts">
import {Login} from "../wailsjs/go/fdu/Fdu";
import router from "@/router";
import { useStore } from "@/stores/store";
import { useI18n } from "vue-i18n";

const { t, availableLocales: languages, locale } = useI18n();

const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};

const onclickMinimise = () => {};

const onclickQuit = () => {};

const store = useStore();
Login([]).then((res) => {
  store.setUser(res);
}).catch((err) => {
  router.push("/login");
})

</script>

<style lang="scss">
@import url("./assets/css/index.css");

header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: #3f51b5;
  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #7986cb;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;
      &:hover,
      &.router-link-exact-active {
        background-color: #3d5afe;
        color: #ffffff;
      }
    }
  }
  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;
    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #7986cb;
      overflow: hidden;
      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;
        &:hover {
          background-color: #7986cb;
          cursor: pointer;
        }
        &.active {
          background-color:  #3d5afe;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }
  }
}
</style>
