<template>
  <div class="flex flex-col items-center" :class="check ? 'checked' : 'unchecked'">
    <div class="nav" @click="check = !check"></div>
    <div class="log-console flex flex-col w-full px-4 bg-white">
      <p v-for="(log, i) in logs" :key="i" class="my-0.5 text-sm">
        {{ log }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, reactive, ref } from "vue";

const check = ref(false)

const logs = reactive([] as string[]);

const updateLog = async (message: string) => {
  logs.push(message)
  if (logs.length > 100) logs.shift()
  await nextTick()
  const logEl = document.querySelector(".log-console")!;
  logEl.scrollTop = logEl.scrollHeight;
}

function connectWS(){
  const socket = new WebSocket("ws://localhost:18964/ws");
  socket.onmessage = (event) => {
    updateLog(event.data)
  };
}
onMounted(() => {
  connectWS()
});

</script>

<style scoped>
.nav {
    position: relative;
    height: 1rem;
    width: 6rem;
    margin-bottom: -1px;
    cursor: pointer;
}

.nav::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: white;
    border: 1px solid grey;
    border-bottom: none;
    border-radius: 1rem 1rem 0 0;
}

.nav::after {
    position: absolute;
    content: '';
    left: 2.5rem;
    bottom: 0.2rem;
    width: 0;
    height: 0;
    border: .5rem solid transparent;
    border-bottom-color: gray;
    transform: rotate(0);
    transition: all 0.5s;
}

.checked .nav::after {
    bottom: -0.5rem;
    transform: rotate(180deg);
}

.log-console {
    height: 0;
    overflow-y: hidden;
    border-top: grey 1px solid;
    transition: height 0.5s;
}

.checked .log-console {
    height: 12.25rem;
    overflow-y: auto;
}

.log-console::-webkit-scrollbar {
    width: 6px;
}

.log-console::-webkit-scrollbar-track {
    background: white;
    border-radius: 2px;
}

.log-console::-webkit-scrollbar-thumb {
    background: rgb(210, 210, 210);
    border-radius: 8px;
}

.log-console::-webkit-scrollbar-thumb:hover {
    background: rgb(180, 180, 180);
}

.log-console::-webkit-scrollbar-corner {
    background: #179a16;
}
</style>