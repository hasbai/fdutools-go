<template>
  <main>
    <div class="mb-6 mt-2 text-center">
      <p class="text-lg font-semibold">
        {{ t("grade.gpa") }}<span class="text-indigo-600">{{ gpa.toFixed(2) }}</span>
      </p>
    </div>
    <table class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-50">
      <tr>
        <th v-for="i in ['semester', 'name', 'grade']"
            class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
          {{t("grade." + i)}}
        </th>
      </tr>
      </thead>
      <tbody class="bg-white divide-y divide-gray-200">
      <tr v-for="(item, i) in paginatedData" :key="i">
        <td class="py-4 px-6 whitespace-nowrap">
          <div class="text-sm text-gray-900">{{item.year}} {{item.semester}}</div>
        </td>
        <td class="py-4 px-6 whitespace-nowrap">
          <div class="text-sm text-gray-900">{{ item.name }}</div>
        </td>
        <td class="py-4 px-6 whitespace-nowrap">
            <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
              {{ item.grade }}
            </span>
        </td>
      </tr>
      </tbody>
    </table>
    <!-- Pagination -->
    <div class="flex justify-center mt-4">
      <nav class="inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <a
          v-if="currentPage > 1"
          @click="changePage(currentPage - 1)"
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
        >
          Previous
        </a>
        <template v-for="i in totalPages" :key="i">
          <a
            @click="changePage(i)"
            :class="{
              'z-10 bg-indigo-50 border-indigo-500 text-indigo-600': currentPage === i,
              'border-gray-300 text-gray-500 hover:bg-gray-50': currentPage !== i
            }"
            class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
          >
            {{ i }}
          </a>
        </template>
        <a
          v-if="currentPage < totalPages"
          @click="changePage(currentPage + 1)"
          class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
        >
          Next
        </a>
      </nav>
    </div>
  </main>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed } from "vue";
import type { fdu } from "../../wailsjs/go/models";
import { GetGrades } from "../../wailsjs/go/fdu/Fdu";
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const tableData = reactive([] as fdu.Grade[]);
const currentPage = ref(1);
const itemsPerPage = 5
const totalPages = computed(() => Math.ceil(tableData.length / itemsPerPage));
const paginatedData = computed(() => {
  const startIndex = (currentPage.value - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  return tableData.slice(startIndex, endIndex);
});
const changePage = (page: number) => {
  currentPage.value = page;
};

const gpa = ref(0)

onMounted(async () => {
  try {
    const response = await GetGrades();
    tableData.push(...response.grades);
    gpa.value = response.gpa;
  } catch (error) {
    console.error(error);
  }
});

</script>

<style scoped>
nav > a {
  cursor: pointer;
}
</style>