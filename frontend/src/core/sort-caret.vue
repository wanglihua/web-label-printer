<script setup>
defineOptions({
    name: "sort-caret",
});

const props = defineProps({
    prop: String,
    label: String,
    sortProp: String,
    sortOrder: String,
    sortMethod: Function,
    cell: Object,
});

function getClassName() {
    let className = "";
    if (props.prop === props.sortProp) {
        if (props.sortOrder === "ascending") {
            className = "ascending";
        }
        if (props.sortOrder === "descending") {
            className = "descending";
        }
    }

    return className;
}

function getSortMethodOrder() {
    props.cell.store.clearSort();

    if (props.sortProp !== props.prop || props.sortOrder === "") {
        return "ascending";
    }

    // below:  props.sortProp === props.prop

    if (props.sortOrder === "ascending") {
        return "descending";
    }

    if (props.sortOrder === "descending") {
        return "";
    }

    return "";
}
</script>

<template>
    <div style="cursor: pointer;" :class="getClassName()" @click="() => sortMethod({ prop: props.prop, order: getSortMethodOrder() })">
        <span>{{ label }}</span>
        <button type="button" class="caret-wrapper" :aria-label="`Sort by ${label}`">
            <i class="sort-caret ascending"></i>
            <i class="sort-caret descending"></i>
        </button>
    </div>

</template>

<style scoped>

</style>
