from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# Headless mode
options = Options()
options.headless = True
options.add_argument("--window-size=1920,1200")
driver = webdriver.Chrome(options=options)

def scroll_to_bottom(driver):
    driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")

def scroll_to_top(driver):
    driver.execute_script("window.scrollTo(0, 0);")

# driver = webdriver.Chrome()
# devpost: for long hackathons, publicise at least a month before deadline
url = "https://devpost.com/hackathons?challenge_type[]=online&status[]=upcoming&status[]=open"
driver.get(url)
target_found = False

while not target_found:
    try:
        target_div = WebDriverWait(driver, 3).until(
            EC.presence_of_element_located((By.XPATH, "//div[contains(@class, 'text-center')]/p[text()='End of results']"))
        )
        
        print("Target element found. Stopping scrolling.")
        target_found = True
    except:
        # due to lazy loading nature, need to scroll down to load the other results
        # odd behaviour of not loading when scrolled to bottom unless we scroll up and then down again
        # hence scroll_to_top is needed
        scroll_to_bottom(driver)
        scroll_to_top(driver)
       

titles = driver.find_elements(By.XPATH, "//div[@class='content']/h3")

for title in titles:
    print(title.text)

driver.quit()