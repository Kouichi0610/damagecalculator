<template>
  <div class="damageInfo">
    {{ result.Target }}
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
import { DefendersResult } from '../../target/store/defenderDamages'

@Component
export default class DamageInfo extends Vue {
  @Prop() private result!: DefendersResult;
}

</script>

<style scoped>
</style>
