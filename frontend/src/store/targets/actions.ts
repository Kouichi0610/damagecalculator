import { ActionTree } from 'vuex';
import axios from 'axios';
import { TargetsState, Species, Filter } from './types';
import { RootState } from '../types';

class SpeciesImpl implements Species {
    name: string;
    types: string;
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;
    constructor (d: any) {
      this.name = d.name;
      this.types = d.types;
      this.hp = d.hp;
      this.attack = d.attack;
      this.defense = d.defense;
      this.spAttack = d.sp_attack;
      this.spDefense = d.sp_defense;
      this.speed = d.speed;
    }
}

export const actions: ActionTree<TargetsState, RootState> = {
    getCandidates: ({ commit }, filter: Filter) => {
        console.log('Get Candidates...');
        axios.get('filtered_list', {
            params: {
              Types: filter.types,
              Total: filter.total,
            }
          })
          .then((response) => {
            let json = JSON.stringify(response.data);
            let sp = JSON.parse(json);
            let species: Species[] = new Array();
            for (var i = 0; i < sp.length; i++) {
              let s = new SpeciesImpl(sp[i]);
              species.push(s);
            }
            commit('setCandidates', species);
          })
          .catch((e) => {
              console.log('failed:' + e);
          });
        return true;
    },

    /*
    async getlist({commit},  types:string, total:number) {
        await axios.get('filtered_list', {
            params: {
              Types: 'ドラゴン',
              Total: '530',
            }
          })
          .then((response) => {
            let json = JSON.stringify(response.data);
            console.log('json:' + json)
            alert('success : ' + response.data)
          })
          .catch((e) => {
            alert('failed ' + e)
            commit('profileError');
          })
        }
    */
}
