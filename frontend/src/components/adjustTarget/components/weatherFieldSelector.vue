<template>
  <div class="weather-fields">
    <div class="row mb-1">
      <b-dropdown class="col-2" id="weather-dropdown" text="天候">
        <b-dropdown-item v-for="item in weathers" :key="item.id" @click="setCurrentWeather(item)">{{ item.name }} {{ item.description }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-2">{{ currentWeather.name }}</div>
      <b-dropdown class="col-2" id="field-dropdown" text="フィールド">
        <b-dropdown-item v-for="item in fields" :key="item.id" @click="setCurrentField(item)">{{ item.name }} {{ item.description }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-2">{{ currentField.name }}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { Weather, Field } from '../../../store/weatherFields/types'

const namespace: string = "weatherFieldsState";

/*
  天候、フィールド選択
*/
@Component
export default class WeatherFieldSelector extends Vue {
  @Action('getWeatherFields', { namespace })
  private getWeatherFields!: () => void;

  @Getter('isInitialized', { namespace })
  private isInitialized!: boolean;
  @Getter('weathers', { namespace })
  private weathers!: Weather[];
  @Getter('fields', { namespace })
  private fields!: Field[];
  @Getter('currentWeather', { namespace })
  private currentWeather!: Weather;
  @Getter('currentField', { namespace })
  private currentField!: Field;
  @Mutation('setCurrentWeather', { namespace })
  private setCurrentWeather!: (weather: Weather) => void;
  @Mutation('setCurrentField', { namespace })
  private setCurrentField!: (field: Field) => void;

  created() {
    if (this.isInitialized) return;
    this.getWeatherFields();
  }
}
</script>
<style scoped>
</style>
