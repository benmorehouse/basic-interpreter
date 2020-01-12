import { newE2EPage } from '@stencil/core/testing';

describe('main-frame', () => {
  it('renders', async () => {
    const page = await newE2EPage();
    await page.setContent('<main-frame></main-frame>');

    const element = await page.find('main-frame');
    expect(element).toHaveClass('hydrated');
  });
});
