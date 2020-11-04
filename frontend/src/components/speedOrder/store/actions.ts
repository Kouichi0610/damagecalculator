import { ActionTree } from 'vuex'
import axios from 'axios';
import { RootState } from '@/store/types';
import { SpeedOrderState, SpeedInfo, SpeedCorrector } from './types';

class SpeedInfoImpl implements SpeedInfo {
  info: string;
  speed: number;
  constructor(info: string, speed: number) {
    this.info = info;
    this.speed = speed;
  }
}

export const actions: ActionTree<SpeedOrderState, RootState> = {
  // 特性の影響を取得
  getAbilityEffect: ({commit}, ability: string) => {
    axios.get('ability_owner_speed?Ability=' + ability)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let a = JSON.parse(json);
      commit('abilityOwner', new SpeedCorrector(a.comment, a.rank, a.magnification));
    })
    .catch((e) => {
      console.log('failed:' + e)
    });

    axios.get('ability_other_speed?Ability=' + ability)
    .then((response) => {
      let json = JSON.stringify(response.data);
      let a = JSON.parse(json);
      commit('abilityOther', new SpeedCorrector(a.comment, a.rank, a.magnification));
    })
    .catch((e) => {
      console.log('failed:' + e)
    });
  },
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
