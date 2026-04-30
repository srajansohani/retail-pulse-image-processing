import requests
import time
import sys

BASE_URL = "http://localhost:8080/api"

def submit_job():
    print("🚀 Submitting image processing job...")
    payload = {
        "count": 2,
        "visits": [
            {
                "store_id": "RP00001",
                "image_url": [
                    "https://images.unsplash.com/photo-1541963463532-d68292c34b19?w=500&auto=format&fit=crop&q=60",
                    "https://images.unsplash.com/photo-1503023345310-bd7c1de61c7d?w=500&auto=format&fit=crop&q=60"
                ],
                "visit_time": "2023-10-27T10:00:00Z"
            },
            {
                "store_id": "RP00002",
                "image_url": [
                    "https://images.unsplash.com/photo-1508921912186-1d1a45ebb3c1?w=500&auto=format&fit=crop&q=60"
                ],
                "visit_time": "2023-10-27T11:30:00Z"
            }
        ]
    }
    # Wait, the count in payload is for total VISITS, but my code checks payload.Count == len(payload.Visits)
    # Let's double check job_handler.go line 26: if payload.Count != len(payload.Visits)
    
    payload["count"] = len(payload["visits"])

    try:
        response = requests.post(f"{BASE_URL}/submit/", json=payload)
        response.raise_for_status()
        job_id = response.json().get("job_id")
        print(f"✅ Job submitted successfully! Job ID: {job_id}")
        return job_id
    except Exception as e:
        print(f"❌ Failed to submit job: {e}")
        sys.exit(1)

def poll_status(job_id):
    print(f"⏳ Polling status for Job ID: {job_id}...")
    while True:
        try:
            response = requests.get(f"{BASE_URL}/status", params={"jobid": job_id})
            response.raise_for_status()
            data = response.json()
            status = data.get("status")
            print(f"🔄 Current Status: {status}")

            if status in ["completed", "failed"]:
                print("\n🏁 Job Finished!")
                print("-" * 30)
                if status == "failed":
                    print(f"❌ Errors found:")
                    for error in data.get("error", []):
                        print(f"   - Store {error['store_id']}: {error['error']}")
                else:
                    print("🎉 All images processed successfully!")
                print("-" * 30)
                break
            
            time.sleep(2)
        except Exception as e:
            print(f"❌ Error polling status: {e}")
            break

if __name__ == "__main__":
    print("🌟 Retail Pulse Image Processing Demo 🌟")
    job_id = submit_job()
    poll_status(job_id)
