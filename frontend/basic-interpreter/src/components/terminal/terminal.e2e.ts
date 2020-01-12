import { newE2EPage } from '@stencil/core/testing';

describe('main-terminal', () => {
  it('renders', async () => {
    const page = await newE2EPage();
    await page.setContent('<main-terminal></main-terminal>');

    const element = await page.find('main-terminal');
    expect(element).toHaveClass('hydrated');
  });
});
