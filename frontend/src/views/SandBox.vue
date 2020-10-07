<template>
  <div class="sandbox">
    <h1>SandBox</h1>
    <damage-result :min="rateMin" :max="rateMax"></damage-result>
    <individuals></individuals>
    <names v-bind:tag="pokemon" v-bind:name="name" v-bind:candidates="candidates"></names>
    <rank></rank>
    <b-form inline>
      <base-points></base-points>
      <base-points></base-points>
    </b-form>
  </div>
</template>

<script>
import Names from '../components/names.vue'
import Individuals from '../components/individuals.vue'
import BasePoints from '../components/basePoints.vue'
import Rank from '../components/rank.vue'
import DamageResult from '../components/damageResult.vue'

export default {
  name: 'SandBox',
  components: {
    'names': Names,
    'individuals': Individuals,
    'base-points': BasePoints,
    'rank': Rank,
    'damage-result': DamageResult,
  },
  data: function() {
    return {
      rateMin: '16.7',
      rateMax: '25.4',
      pokemon: 'ポケモン',
      name: 'ピカチュウ',
      candidates: ['アーボック', 'フシギバナ', 'ギャラドス', 'リザードン', 'カメックス', 'ピカチュウ', 'カイリュー', 'ドサイドン', 'ミュウ', 'ホルード', 'ランクルス', 'エルレイド', 'バタフリー'],
    }
  },
  watch: {
    name: function(next, prev) {
      console.log('Prev' + prev + ' Next:' + next);
    }

  }
}
</script>

<style scoped>
</style>