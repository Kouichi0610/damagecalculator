// 調整するポケモンを選択
// 現在の振り方
// TODO:変更時に保存確認
<template>
  <div class="current">
    <h1>現在のポケモン</h1>
    <target-select></target-select>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TargetSelect from '../components/targetSelect.vue'

@Component({
  components: {
    TargetSelect,
  },
})
export default class Current extends Vue {
}
</script>

<style scoped>

</style>
