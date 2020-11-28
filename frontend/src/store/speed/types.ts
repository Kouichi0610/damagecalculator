import axios from 'axios'
import { TargetCondition } from '../target/targetCondition'

export interface SpeedState {
  target: string,
  level: number;
  otherList: SpeedInfo[];
  targetSpeed: number; // とくせい、もちもの込みの素早さ
  trickRoom: boolean;
}

// すばやさ取得
export class GetSpeed {
  getSpeed(target: TargetCondition): Promise<number> {
    return new Promise((resolve, reject) => {
      axios.get('get_speed', {
        params: {
          Name: target.target,
          Level: target.level,
          Individuals: target.individuals.type(),
          BasePoint: target.basePoints.speed(),
          Ability: target.ability,
          Nature: target.nature.name,
          Item: target.item.name,
          Condition: target.condition,
          Weather: target.weather.name,
          Field: target.field.name,
          }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let result = JSON.parse(json);
        resolve(result.speed)
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(0);
      });
    });
  }
}

// 仮想敵　一覧 取得
export class SpeedList {
  speedList(level: number): Promise<SpeedInfo[]> {
    return new Promise((resolve, reject) => {
      axios.get('speed_list?Level=' + level)
      .then((response) => {
        let json = JSON.stringify(response.data);
        let list = JSON.parse(json);
        let res: SpeedInfo[] = [];
        for (var i = 0; i < list.length; i++) {
          res.push(new SpeedInfo(i, list[i].info, list[i].speed, false));
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject([]);
      });
    });
  }

}

export class SpeedInfo {
  readonly index: number;
  readonly info: string;
  readonly speed: number;
  readonly target: boolean;

  constructor(index: number, info: string, speed: number, target: boolean) {
    this.index = index;
    this.info = info;
    this.speed = speed;
    this.target = target;
  }
}

// 
export class SpeedRanking {
  ranking(target: string, targetSpeed: number, trickRoom: boolean, others: SpeedInfo[]): SpeedInfo[] {
    let empty: SpeedInfo[] = [];
    let res: SpeedInfo[] = empty.concat(others);
    res.push(new SpeedInfo(-1, target, targetSpeed, true));
    let decending = function(a: SpeedInfo, b: SpeedInfo) {
      if (a.speed > b.speed) return -1;
      if (a.speed < b.speed) return 1;
      return 0;
    }
    let ascending = function(a: SpeedInfo, b: SpeedInfo) {
      if (a.speed < b.speed) return -1;
      if (a.speed > b.speed) return 1;
      return 0;
    }
    let order = trickRoom ? ascending : decending;
    res.sort(order);
    return res;
  }
}
