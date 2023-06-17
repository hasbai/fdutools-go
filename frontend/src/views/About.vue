<script setup lang="ts">
import { useI18n } from "vue-i18n";

import gif from "@/assets/images/comeon.gif";

const { t } = useI18n();
</script>

<template>
  <main class="about">
    <div class="title mt-4 mb-8">{{ t("aboutpage.title") }}</div>
    <div class="content flex flex-row justify-around">
      <img :src="gif" alt="Gif" />
      <ul class="info flex flex-col">
        <li class="info-item">
          <div class="name">{{ t("aboutpage.project-repository") }}</div>
          <a class="link" href="https://github.com/hasbai/fdutools-go"
            >https://github.com/hasbai/fdutools-go</a
          >
        </li>
        <li class="info-item">
          <div class="name">{{ t("aboutpage.author") }}</div>
          <a class="link" href="https://github.com/hasbai">{{
            t("aboutpage.name")
          }}</a>
        </li>
      </ul>
    </div>
    <div class="thank">{{ t("aboutpage.thanks") }}</div>
  </main>
</template>

<style scoped lang="scss">
.about {
  .title {
    font-size: 38px;
    color: #a150b5;
    text-align: center;
  }
  .content {
    img {
      max-width: 30%;
    }
    .info {
      text-align: left;
      .info-item {
        margin-bottom: 10px;
        .name {
          line-height: 40px;
          font-size: 28px;
          color: #6d6363;
        }
        .link {
          line-height: 30px;
          font-size: 20px;
          color: #5f6c86;
        }
      }
    }
  }

  .thank {
    height: 68px;
    line-height: 68px;
    margin: 36px auto;
    text-align: center;
    font-size: 40px;
  }
}
</style>
