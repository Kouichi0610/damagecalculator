// TODO:
import {mount} from '@vue/test-utils'
import { shallowMount } from '@vue/test-utils'
import {sum} from './sum'

describe('sum sample', (): void => {
  test('1+1 = 2', (): void => {
    let response: number = sum(1, 2);
    expect(response).toBe(3);
  });
})

