

Assumptions :
1. Assuming one item can be added as part of the deal
2. Quantity of the items can be many
3. Time has been represented in the int64 (epoch time style)

Happy Flows :
1. Add users to the system using the GetUser API
2. Add deals to the system using the GetDeal API
3. End a deal using the EndDeal API
4. Update the deal using the UpdateDeal API
5. Claim a deal using the ClaimDeal API

Validation Checks to claim a deal :
1. A single user can only claim the deal once
2. A user cannot claim the deal once it has been ended
3. A user cannot claim the deal if all items have finished
4. A user cannot claim a deal if time has expired for the deal

// Expired time : 1691426039
// Valid time : 1754584439