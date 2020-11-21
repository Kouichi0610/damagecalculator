<template>
  <div class="WeatherFields">
    <div class="row mb-1">
      <div class="col-1"></div>
      <b-dropdown class="col-1" id="weather-dropdown" text="天候">
        <b-dropdown-item v-for="item in weathers" :key="item" @click="setCurrentWeather(item)">{{ item }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{ currentWeather }}</div>
      <b-dropdown class="col-1" id="field-dropdown" text="フィールド">
        <b-dropdown-item v-for="item in fields" :key="item" @click="setCurrentField(item)">{{ item }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{ currentField }}</div>
    </div>
  </div>
</template>

<script lang="ts">
/*
  天候、フィールド
*/
import { Vue, Prop, Component, Watch } from "vue-property-decorator";
import { Getter, Action, Mutation } from 'vuex-class';
import { WeatherFieldsState } from '../../store/weatherFields/types';

const namespace: string = 'weatherFields';

@Component
export default class WeatherFields extends Vue {
  @Action('getWeatherFields', { namespace })
  private getWeatherFields!: () => void;
  @Getter('isInitialized', { namespace })
  private isInitialized!: boolean;
  @Getter('weathers', { namespace })
  private weathers!: string[];
  @Getter('fields', { namespace })
  private fields!: string[];
  @Getter('currentWeather', { namespace })
  private currentWeather!: string;
  @Getter('currentField', { namespace })
  private currentField!: string;
  @Mutation('setCurrentWeather', { namespace })
  private setCurrentWeather!: (current: string) => void;
  @Mutation('setCurrentField', { namespace })
  private setCurrentField!: (current: string) => void;

  created() {
    this.weatherChanged();
    this.fieldChanged();
    if (this.isInitialized) {
      return;
    }
    this.getWeatherFields();
  }

  @Watch('currentWeather')
  weatherChanged() {
    if (this.currentWeather.length == 0) return;
    this.$emit('weather', this.currentWeather);
  }
  @Watch('currentField')
  fieldChanged() {
    if (this.currentField.length == 0) return;
    this.$emit('field', this.currentField);
  }
}
</script>

<style scoped>
</style>
