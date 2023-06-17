<template>
    <div class="hello-world" v-text="msg"></div>
</template>


<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<style scoped>
.hello-world {
  height: 76px;
  line-height: 38px;
  margin: 16px 150px;
  text-align: center;
  font-size: 26px;
}
</style>
