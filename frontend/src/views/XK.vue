<template>
  <main>
    <table class="min-w-full text-center divide-y divide-gray-200 table-fixed">
      <thead class="bg-gray-50">
        <tr>
          <th class="w-16"></th>
          <th
            v-for="i in 7"
            class="p-1 w-8 text-xs font-medium text-gray-500 tracking-wider"
          >
            {{ t("xk.col" + (i - 1).toString()).substring(0, 3) }}
          </th>
        </tr>
      </thead>
      <tbody
        class="bg-white divide-x divide-y divide-gray-200 text-sm text-gray-900"
      >
        <tr v-for="(row, i) in tableData" :key="i" class="h-8">
          <td>{{ timeTable[i] }}</td>
          <td
            v-for="(col, j) in row"
            :key="j"
            class="p-1 whitespace-nowrap"
            @click="onClick"
          >
            {{ col }}
          </td>
        </tr>
      </tbody>
    </table>
    <XkSearch :show="showSearch" />
    <XkOption />
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed } from "vue";
import { xk } from "../../wailsjs/go/models";
import { QueryCourses, Login } from "../../wailsjs/go/xk/Xk";
import { useI18n } from "vue-i18n";
import XkSearch from "@/components/XkSearch.vue";
import XkOption from "@/components/XkOption.vue";
import { useStore } from "@/stores/store";
const { t } = useI18n();

const courses = reactive([] as xk.Course[]);

const tableData = computed(() => {
  const data = reactive([] as string[][]);
  for (let i = 0; i < 14; i++) {
    data.push([]);
    for (let j = 0; j < 7; j++) {
      data[i].push("");
    }
  }
  courses.forEach((course) => {
    const arrange = course.arrangeInfo[0];
    for (let i = arrange.startUnit; i <= arrange.endUnit; i++) {
      data[i - 1][arrange.weekDay] = course.name;
    }
  });
  return data;
});

const timeTable = [
  "08:00-08:45",
  "08:55-09:40",
  "09:55-10:40",
  "10:50-11:35",
  "11:45-12:30",
  "13:30-14:15",
  "14:25-15:10",
  "15:25-16:10",
  "16:20-17:05",
  "17:15-18:00",
  "18:30-19:15",
  "19:25-20:10",
  "20:20-21:05",
  "21:15-22:00",
];

const store = useStore();
const showSearch = ref(false);
const onClick = (e: Event) => {
  const target = e.target! as HTMLElement;
  if (target.innerText === "") {
    showSearch.value = !showSearch.value;
  } else {
    showSearch.value = false;
    store.setCurrentCourse(courses.find((c) => c.name === target.innerText)!);
  }
};

onMounted(async () => {
  try {
    await Login();
    let resp = await QueryCourses(new xk.Query());
    courses.push(...resp);
  } catch (error) {
    console.error(error);
  }
});
</script>

<style scoped></style>
