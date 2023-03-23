<template>

</template>

<script setup lang="ts">
const fn = ({name,age}:{name:string,age:number}) => {
  console.log(name,age)
  return name
}
console.log(fn({name:"sz",age:1}))
</script>

<style scoped>

</style>