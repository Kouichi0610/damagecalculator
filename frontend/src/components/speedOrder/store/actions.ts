import { ActionTree } from 'vuex'
import axios from 'axios';
import { RootState } from '@/store/types';
import { SpeedOrderState, SpeedInfo } from './types';

class SpeedInfoImpl implements SpeedInfo {
  info: string;
  speed: number;
  constructor(info: string, speed: number) {
    this.info = info;
    this.speed = speed;
  }
}

export const actions: ActionTree<SpeedOrderState, RootState> = {
  getOrders: ({commit}, level: number) => {
    axios.get('speed_list?Level=' + level)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let list = JSON.parse(json);

      let speeds: SpeedInfo[] = new Array();
      for (var i = 0; i < list.length; i++) {
        let s = new SpeedInfoImpl(list[i].info, list[i].speed);
        speeds.push(s)
      }
      commit('setInfos', speeds);
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
  },
}
