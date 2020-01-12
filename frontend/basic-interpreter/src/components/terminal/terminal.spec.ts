import { Terminal } from './terminal';

describe('main-terminal', () => {
  it('builds', () => {
    expect(new Terminal()).toBeTruthy();
  });
});
