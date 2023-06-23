<template>
  <div class="relative">
    <div class="absolute w-full h-full flex justify-center items-center">
      <span class="text-cyan-700 font-bold text">{{ percentage }}%</span>
    </div>
    <svg :width="props.size" :height="props.size">
      <circle
        class="bg"
        :stroke-dasharray="length"
        :stroke-offset="progress"
      ></circle>
      <circle
        class="progress"
        :stroke-dasharray="progress + ' ' + inf"
      ></circle>
    </svg>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, nextTick } from "vue";
const props = defineProps({
  percentage: {
    type: Number,
    required: true,
    default: 0,
  },
  size: {
    type: Number,
    default: 100,
  },
});
const inf = 1 << 20;
const strokeWidth = computed(() => Math.round(props.size / 8));
const cx = computed(() => props.size / 2);
const radius = computed(() => props.size / 2 - strokeWidth.value);
const length = computed(() => 2 * Math.PI * radius.value);
const progress = computed(() => (props.percentage / 100) * length.value);
const fontSize = computed(() => props.size / 5 + "px");
onMounted(() => {});
</script>

<style scoped>
svg {
  transform: rotate(-90deg);
}
circle {
  fill: none;
  stroke-width: v-bind(strokeWidth);
  cx: v-bind(cx);
  cy: v-bind(cx);
  r: v-bind(radius);
}
.bg {
  stroke: lightgrey;
}
.progress {
  stroke: rgb(15, 118, 110);
  transition: stroke-dasharray 0.5s ease;
}
.text {
  font-size: v-bind(fontSize);
}
</style>
