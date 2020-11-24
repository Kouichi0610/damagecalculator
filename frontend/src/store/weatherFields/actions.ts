import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { WeatherFieldsState, toWeathers, toFields } from './types'

export const actions: ActionTree<WeatherFieldsState, RootState> = {
  getWeatherFields: ({ commit } ) => {
    axios.get('weather_fields')
    .then((response) => {
      let json = JSON.stringify(response.data);
      let res = JSON.parse(json);

      let weathers = toWeathers(res);
      let fields = toFields(res);

      commit('setWeathers', weathers);
      commit('setFields', fields);
    })
    .catch((e) => {
      console.log('failed:' + e);
    });
  }
}
