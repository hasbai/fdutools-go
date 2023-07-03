<template>
  <Transition name="fade">
    <div
      v-show="props.show"
      class="flex flex-col absolute top-0 right-0 m-2 px-4 py-2 bg-white shadow-lg rounded-md"
    >
      <div class="pt-4 pb-1 flex flex-row">
        <input
          ref="search"
          autofocus
          v-model="text"
          type="text"
          class="flex-1 px-4 py-2 mr-4 text-gray-800 bg-gray-100 border-2 border-gray-200 rounded-md"
          placeholder="Search..."
          @keyup.enter.prevent="performSearch"
        />
        <button
          @click="performSearch"
          class="w-12 p-2 text-white bg-blue-600 rounded-md hover:bg-blue-500"
        >
          🔎
        </button>
      </div>
      <table class="w-full text-center table-fixed divide-y divide-gray-300">
        <thead class="font-semibold">
          <tr>
            <th class="w-1/4">{{ t("xk.name") }}</th>
            <th>{{ t("xk.teacher") }}</th>
            <th>{{ t("xk.credit") }}</th>
            <th>{{ t("xk.campus") }}</th>
            <th>{{ t("xk.exam") }}</th>
            <th>{{ t("xk.amount") }}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="course in courses" :key="course.id">
            <td>{{ course.name }}</td>
            <td>{{ course.teachers }}</td>
            <td>{{ course.credits }}</td>
            <td>{{ course.campusName }}</td>
            <td>{{ course.examFormName }}</td>
            <td>{{ course.amount.sc }} / {{ course.amount.lc }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { nextTick, onMounted, onUpdated, reactive, ref, watch } from "vue";
import { QueryCourses } from "../../wailsjs/go/xk/Xk";
import { xk } from "../../wailsjs/go/models";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const props = defineProps({
  show: Boolean,
});
const text = ref("");
const courses = reactive([] as xk.Course[]);

const isAscii = (str: string) => /^[\x00-\x7F]+$/.test(str);
const performSearch = async () => {
  const query = new xk.Query();
  if (text.value.includes(".")) {
    query.lessonNo = text.value;
  } else if (isAscii(text.value)) {
    query.courseCode = text.value;
  } else {
    query.courseName = text.value;
  }
  try {
    const res = await QueryCourses(query);
    console.log(res);
    courses.length = 0;
    courses.push(...res);
  } catch (e) {
    console.error(e);
  }
};
const search = ref<HTMLInputElement | null>(null);
watch(props, () => {
  if (props.show) {
    nextTick(() => {
      search.value?.focus();
    });
  }
});
</script>

<style scoped>
td,
th {
  padding: 0.5rem;
}

table {
  min-height: 200px;
}
</style>
