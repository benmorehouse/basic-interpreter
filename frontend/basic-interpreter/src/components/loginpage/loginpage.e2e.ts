import { newE2EPage } from '@stencil/core/testing';

describe('main-loginpage', () => {
  it('renders', async () => {
    const page = await newE2EPage();
    await page.setContent('<main-loginpage></main-loginpage>');

    const element = await page.find('main-loginpage');
    expect(element).toHaveClass('hydrated');
  });
});
