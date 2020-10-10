// ランク補正
<template>
      <div class="rank">
        <b-dropdown id="dropdown-1" text="ランク" class="m-md-2">
            <b-dropdown-item v-for="(item, key) in ranks" :key="key" @click="rankValue = item">{{ item }}</b-dropdown-item>
        </b-dropdown>
        {{ rankValue }}
    </div>
</template>

<script>
export default {
    name: 'Rank',
    components: {
    },
    created: function() {
    },
    data: function() {
        return {
            rankValue: 0,
            ranks: [-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6],
        }
    },
    watch: {
    }
}
</script>

<style scoped>

</style>
