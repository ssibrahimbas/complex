<template>
  <div class="fib-page">
  <form @submit.prevent="handleSubmit">
    <label>
      Enter your index:
    </label>
    <input v-model="index" />
    <button>Submit</button>
  </form>

  <h3>Indexes I have seen:</h3>
    {{seenIndexes.map(({values}) => values).join(", ")}}
  <h3>Calculated Values:</h3>
    <div v-for="(item, key) in values" :key="key">
      For index {{key}} I calculated {{item}}
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from "vue";
import {http} from "../plugins/axios.js";

const seenIndexes = ref([]);
const values = ref({});
const index = ref('');

async function fetchValues() {
  const _values = await http.get("/api/values/current");
  values.value = _values.data.data
}

async function fetchIndexes() {
  const _seenIndexes = await http.get("/api/values/all");
  seenIndexes.value = _seenIndexes.data.data
}

async function handleSubmit() {
    console.log('submit')
  await http.post("/api/values", {
    Values: +index.value
  })
  index.value = ''
}

onMounted(() => {
  fetchValues()
  fetchIndexes()
})


</script>

<style scoped>

</style>