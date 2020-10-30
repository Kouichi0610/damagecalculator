import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types'
import { TargetState, Species, StatsPatternArgs } from './types'

export const actions: ActionTree<TargetState, RootState> = {
  getStatsPattern: ({commit}, args: StatsPatternArgs) => {
    axios.get('stats_pattern', {
      params: {
        Level: args.level,
        Name: args.name,
        Nature: args.nature,
        HP: args.hp,
        Attack: args.at,
        Defense: args.df,
        SpAttack: args.sa,
        SpDefense: args.sd,
        Speed: args.sp,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      let res = JSON.parse(json);
      // TODO:getter
      // TODO:args多すぎ＆res多すぎ
      commit('setStatsPattern', res);
    })
    .catch((e) => {
        console.log('failed:' + e);
    });
  },
  getSpecies: ({commit}, name: string) => {
    axios.get('get_species', {
      params: {
        Name: name,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      let sp = JSON.parse(json);
      commit('setName', sp.Name);
      commit('setTypes', sp.Types);
      commit('setWeight', sp.Weight);
      let species: Species = {hp: sp.Species[0], at: sp.Species[1], df: sp.Species[2], sa: sp.Species[3], sd: sp.Species[4], sp: sp.Species[5]};
      commit('setSpecies', species);
      commit('setAbilities', sp.Abilities);
      return true;
    })
    .catch((e) => {
        console.log('failed:' + e);
        return false;
    });
  }
}

