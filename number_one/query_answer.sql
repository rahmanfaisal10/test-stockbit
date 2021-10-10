SELECT 
    u.ID, 
    u.UserName, 
    u2.UserName as ParentUserName 
FROM `USER` u 
LEFT JOIN `USER` u2 ON u.Parent = u2.ID;