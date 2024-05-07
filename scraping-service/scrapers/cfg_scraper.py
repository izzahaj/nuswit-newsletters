from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# uses javascript
# view events in list view (use selenium to click the button)
# only shows upcoming events

# Headless mode

# TODO: find a way to abstract away all the common boilerplate code
options = Options()
options.headless = True
options.add_argument("--window-size=1920,1200")
driver = webdriver.Chrome(options=options)

url = "https://nus.edu.sg/cfg/events"
driver.get(url)

# wait to load
# WebDriverWait(driver, 2).until(EC.)

driver.quit()
