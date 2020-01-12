import { newE2EPage } from '@stencil/core/testing';

describe('main-aboutpage', () => {
  it('renders', async () => {
    const page = await newE2EPage();
    await page.setContent('<main-aboutpage></main-aboutpage>');

    const element = await page.find('main-aboutpage');
    expect(element).toHaveClass('hydrated');
  });
});
