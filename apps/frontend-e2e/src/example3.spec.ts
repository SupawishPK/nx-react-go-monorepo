import { test, expect } from '@playwright/test';

test('has title', async ({ page }) => {
  await page.goto('/');

  // Expect h1 to contain a substring.
  expect(await page.locator('h1').innerText()).toContain('Welcome');
});

test('click Add UI library', async ({ page }) => {
  await page.goto('/');
  await page.click('text=Add UI library');
  await expect(page.getByText('nx g @nx/next:library ui')).toBeVisible();
});
