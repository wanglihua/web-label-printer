<script setup>
defineOptions({
    name: "data-dump",
});

const props = defineProps({
    data: Object,
});

// When the environment cannot print output, use this component to output data
console.log(props.data);

</script>

<template></template>

<style scoped></style>
