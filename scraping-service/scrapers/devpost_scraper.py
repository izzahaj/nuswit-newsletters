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

# devpost: for long hackathons, publicise at least a month before deadline
url = "https://devpost.com/hackathons?challenge_type[]=online&status[]=upcoming&status[]=open"
driver.get(url)
target_found = False

while not target_found:
    try:
        target_div = WebDriverWait(driver, 2).until(
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
       
title_elements = driver.find_elements(By.XPATH, "//div[@class='content']/h3")
location_elements = driver.find_elements(By.XPATH, "//div[@class='content']/div/div/div/div[@class='info-with-icon']/div[@class='info']/span")
link_elements = driver.find_elements(By.XPATH, "//a[@class='flex-row tile-anchor']")
titles = [title_element.text for title_element in title_elements]
locations = [location_element.text for location_element in location_elements]
links = [link_element.get_attribute("href") for link_element in link_elements]

start_dates = []
end_dates = []

# parse links and access schedule
for i, link in enumerate(links):
    parts = link.split("/")
    domain = "/".join(parts[:3])
    links[i] = domain
    schedule_url = f"{domain}/details/dates"

    driver.get(schedule_url)

    submissions_td = driver.find_element(By.XPATH, "//td[contains(text(), 'Submissions')]")
    start_td = submissions_td.find_element(By.XPATH, "./following-sibling::td")
    start_dates.append(start_td.get_attribute("data-iso-date"))
    end_td = start_td.find_element(By.XPATH, "./following-sibling::td")
    end_dates.append(end_td.get_attribute("data-iso-date"))

for title, location, start_date, end_date, link in zip(titles, locations, start_dates, end_dates, links):
    print(title)
    print(location)
    print(start_date)
    print(end_date)
    print(link)
    
driver.quit()
